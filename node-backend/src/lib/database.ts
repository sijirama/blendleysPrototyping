import { PrismaClient } from '@prisma/client';

export const prisma = new PrismaClient();

export async function DatabaseConnect() {
    try {
        await prisma.$connect();
        console.log('Database Connection established ');
        return prisma;
    } catch (e) {
        console.error(e);
        await prisma.$disconnect();
        process.exit(1);
    }
}
