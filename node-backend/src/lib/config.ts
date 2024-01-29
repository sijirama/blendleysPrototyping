import dotenv from 'dotenv';

dotenv.config();

export const CONFIG = {
    DATABASE_URL: String(process.env.DATABASE_URL),
    PORT: process.env.PORT || 3000,
    SALT: Number(process.env.SALT),
    JWT_SECRET: String(process.env.JWT_SECRET),
    COOKIE_SECRET: String(process.env.COOKIE_SECRET),
    JWT_TOKEN_NAME: 'token'
};
