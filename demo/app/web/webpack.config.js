const path = require('path');

module.exports = {
    entry: './script.js', // Specify the entry point of your application
    output: {
        path: path.resolve(__dirname, 'dist'), // Specify the output directory for bundled files
        filename: 'bundle.js', // Specify the name of the bundled file
    },
};