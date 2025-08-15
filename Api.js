const fetch = require('node-fetch');

const url = 'https://evilinsult.com/generate_insult.php?lang=en&type=json';

fetch(url)
    .then(response => response.json())
    .then(data => console.log(data))
    .catch(error => console.error('Error:', error));