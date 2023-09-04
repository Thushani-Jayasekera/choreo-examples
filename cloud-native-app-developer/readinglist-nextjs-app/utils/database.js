import mysql from 'mysql2/promise';

let isConnected = false;
let connection = null;

export const connectToDB = async () => {
    isConnected = false;
    if (isConnected) {
        console.log('Database is already connected');
        return connection;
    }

    try {
        connection = await mysql.createConnection({
            host: process.env.MYSQL_HOST,
            user: process.env.MYSQL_USER,
            password: process.env.MYSQL_PASSWORD,
            database: 'books'
        });

        isConnected = true;
        console.log('Connected to the database');
        return connection;
    } catch (error) {
        console.log(error);
    }
};
