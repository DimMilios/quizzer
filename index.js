const fs = require('fs');

const file = fs.readFileSync('./tasklist_out.txt', 'utf8');

function convert(data) {
    try {
        return JSON.parse(JSON.stringify(data));
    } catch (err) {
        console.error("Failed to parse json data", err);
    }
}

console.log(convert(file));
