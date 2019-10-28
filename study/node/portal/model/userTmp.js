'use strict';
const loader = require('./_sequelize-loader');
const Sequelize = loader.Sequelize;

const UserTmp = loader.database.define('user_tmp', {
    email: {
        type: Sequelize.STRING,
        primaryKey: true
    },
    username: {
        type: Sequelize.STRING,
        allowNull: false
    },
    password: {
        type: Sequelize.STRING,
        allowNull: false
    },
    token: {
        type: Sequelize.STRING,
        allowNull: false,
        unique: true
    }
}, {
    freezeTableName: true,
    timestamps: false
});

module.exports = UserTmp;