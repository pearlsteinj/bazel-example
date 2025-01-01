const createError = require('http-errors');
const express = require('express');
const router = express.Router();
const path = require('path');
const cookieParser = require('cookie-parser');
const logger = require('morgan');
const upload = require('multer')();
const session = require('express-session');


const app = express();
app.use(session({secret: "Your secret key", resave:true, saveUninitialized:false, cookie: { secure: false } }));

app.use(logger('dev'));
app.use(express.json());
app.use(express.urlencoded({ extended: false }));
app.use(cookieParser());
app.use(upload.array());
app.use(express.static(path.join(__dirname, 'public')));

router.get('/', function(req, res) {
    res.send("Hello, World!");
});

app.use(router);

module.exports = app;
