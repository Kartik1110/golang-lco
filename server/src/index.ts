import express, { Request, Response } from 'express';
import bodyParser from 'body-parser';
import multer from 'multer';

const app = express();
const port = 3000;

// Middleware to parse JSON body
app.use(bodyParser.json());

// Middleware for form data
const upload = multer();

// 1. /get route - a simple GET request that returns a response
app.get('/get', (req: Request, res: Response) => {
    res.send('This is a GET route!');
});

// 2. /post route - a POST route that accepts JSON data and returns a response
app.post('/post-json', (req: Request, res: Response) => {
    const jsonData = req.body;
    res.json({
        message: 'JSON data received successfully!',
        data: jsonData
    });
});

// 3. /post route - a POST route that accepts form data and returns a response
app.post('/post-form', upload.none(), (req: Request, res: Response) => {
    const formData = req.body;
    res.json({
        message: 'Form data received successfully!',
        data: formData
    });
});

// Start the server
app.listen(port, () => {
    console.log(`Server is running on http://localhost:${port}`);
});
