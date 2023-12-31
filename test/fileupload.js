const axios = require('axios');
const fs = require('fs');

const file = fs.createReadStream('/home/niujh/下载/mia/data/face/test_20/test/10046.png');
console.log("file read success")
console.log(file)
// const url = 'http://localhost:8080/user/FileUpload';
// const formData = new FormData();
// formData.append('file', file);
// fetch(url, {
//     method: 'POST',
//     body: formData
// })
//     .then(response => {
//         if (response.ok) {
//             console.log('File uploaded successfully');
//         } else {
//             console.error('File upload failed');
//         }
//     })
//     .catch(error => {
//         console.error('Error:', error);
//     });
