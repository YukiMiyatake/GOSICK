'use strict';
const express = require('express');
const router = express.Router();

// パッケージの読み込み
const crypto = require('crypto');
const uuid = require('uuid');

// モデルの読み込み
const UserTmp = require('../models/userTmp');
const User = require('../models/user');
const UserAuth = require('../models/userAuth');

// '~/auth/email/:token'にGETアクセスが来たときの処理
router.get('/email/:token', function (req, res, next) {
    const token = hashing(req.params.token);
    UserTmp.findOne({
        where: {
            token: token
        }
    }).then(user => {
        if (user) {
            const userId = uuid.v4();
            const username = user.username;
            const password = user.password;
            const email = user.email;
            User.create({
                user_id: userId,
                username: username
            }).then(() => {
                UserAuth.create({
                    username: username,
                    user_id: userId,
                    password: password,
                    email: email
                }).then(() => {
                    UserTmp.destroy({
                        where: {
                            email: email
                        }
                    }).then(() => {
                        res.render('auth', { title: '認証成功' });
                    });
                });
            });
        } else {
            res.render('auth', { title: '認証失敗' });
        }
    });
});

// ハッシュ化関数
function hashing(data) {
    const shasum = crypto.createHash('sha1');
    shasum.update(data);
    let hash = shasum.digest('hex');
    return hash;
}

module.exports = router;