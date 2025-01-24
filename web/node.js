const express = require('express');
const cors = require('cors');
const axios = require('axios');
const fs = require('fs');
const https = require('https');

// SSL uchun sertifikatlar manzili
const sslOptions = {
  key: fs.readFileSync('/etc/letsencrypt/live/transfer-travel.uz/privkey.pem'),
  cert: fs.readFileSync('/etc/letsencrypt/live/transfer-travel.uz/fullchain.pem')
};

const app = express();
app.use(express.json()); // JSON formatdagi so'rovlarni qabul qilish uchun

// CORS konfiguratsiyasi
app.use(cors({
  origin: ['https://admin.transfer-travel.uz', 'https://transfer-travel.uz'],
  credentials: true
}));

// /client marshrutini qo'shish
app.post('/client', async (req, res) => {
  try {
    console.log('Kelayotgan ma\'lumotlar:', req.body);

    // Backendga ma'lumot yuborish
    const response = await axios.post('https://1021-217-30-173-218.ngrok-free.app/web/', req.body, {
      headers: {
        'Content-Type': 'application/json',
        'X-CSRFToken': req.headers['x-csrftoken'] || '' // CSRF token yuborilishi uchun
      }
    });

    // Backenddan kelayotgan javobni qaytarish
    res.status(response.status).send(response.data);
  } catch (error) {
    console.error('So\'rovni qayta ishlashda xato yuz berdi:', error.message);
    // Xatolik haqida ma'lumotni aniqroq qaytarish
    res.status(error.response ? error.response.status : 500).send({
      message: 'So\'rovni qayta ishlashda xato yuz berdi.',
      error: error.response ? error.response.data : 'Noma\'lum xato'
    });
  }
});

// HTTPS serverni ishga tushirish
https.createServer(sslOptions, app).listen(3000, () => {
  console.log('Server is running on port 3000');
});
// root@haonhaudvv:/var/www/html/Luxury_UAE_web/js#
