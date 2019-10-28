'use strict';
const Sequelize = require('sequelize');
/*
const sequelize = new Sequelize(
    'postgres://root:root@postgresql/gosickdb',
    { logging: true, operatorsAliases: false, timezone: 'UTC' }
);
*/
const sequelize = new Sequelize(
    'gosickdb', 'root', 'root',
    { host: 'postgresql', dialect: 'postgres', timezone: 'UTC' }
);

module.exports = {
    database: sequelize,
    Sequelize: Sequelize
};