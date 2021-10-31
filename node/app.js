require('dotenv').config()
const { WebClient } = require('@slack/web-api');
const cron = require('node-cron');

const token = process.env.SLACK_TOKEN;
const web = new WebClient(token);




const msg = process.env.NEW_THREAD_MESSAGE;
const channel = process.env.NEW_THREAD_CHANNEL;
const msg_time = process.env.NEW_THREAD_TIME;

// 9時に新スレ立てる    
//cron.schedule('0 0 ${msg_time} * *', async () => {
cron.schedule('1-59 * * * *', async () => {
    const result = await web.chat.postMessage({
        text: msg,
        channel: channel,
    });
});


