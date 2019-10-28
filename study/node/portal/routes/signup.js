'use strict';
const express = require('express');
const router = express.Router();
const crypto = require("crypto");

const nodemailer = require('nodemailer');

// メール送信設定
const transporter = nodemailer.createTransport({
    service: 'Gmail',
    auth: {
        user: 'yourmail@gmail.com', // Gmailアドレス
        pass: 'hogehuga' // Googleアカウントのパスワード
    }
});
const mailOptions = {
    from: 'yourmail info <yourmail@gmail.com>',
    subject: 'アカウントの確認'
};

// モデルの読み込み
const UserTmp = require('../models/userTmp');

// '~/signup'にPOSTアクセスが来たときの処理
router.post('/', function (req, res, next) {
    const email = req.body.email;
    const username = req.body.username;
    const password = hashing(req.body.password);
    const token = crypto.randomBytes(16).toString('hex');
    const hashedToken = hashing(token);
    UserTmp.upsert({
        email: email,
        username: username,
        password: password,
        token: hashedToken
    }).then(() => {
        mailOptions.to = email;
        mailOptions.html = '<p>以下のリンクからアカウントの確認を行ってください｡</p><br>< a href = "localhost:3000/auth/email/' + token + '" > アカウントを確認</a > ';
        transporter.sendMail(mailOptions, (err, info) => {
            if (err) {
                console.log(err);
            } else {
                console.log('Message sent: ' + info.accepted);
                res.render('signup', { title: 'アカウントの確認' });
            }
        });
    });
});

// ハッシュ化関数
function hashing(data) {
    const shasum = crypto.createHash('sha1');
    shasum.update(data);
    let hash = shasum.digest('hex');
    return hash;
}

// '~/signup'にGETアクセスが来たときの処理
router.get('/', function (req, res, next) {
    res.render('signup', { title: 'サインアップ' });
});

module.exports = router;
