export enum UserRole {
    EMPLOYEE,
    MANAGER,
    ADMIN,
}

export type User = {
    id: string;
    name: string;
    role: UserRole;
    area: string;
};

export enum PaymentType {
    CASH,
    CARD,
}

export enum ClientType {
    PERSONAL,
    MANAGED,
}

export type Purchase = {
    // ID
    id: string;
    // Тип договора
    paymentType: PaymentType;
    // Время создания
    createdAt: Date;
    // Создатель заявки
    creator: User;
    // Поставщик
    supplier: string;
    // Объем (в литрах)
    volumeInLiters: number;
    // Цена (за литр)
    pricePerLiter: number;
    // Стоимость
    totalPrice: number;
    // Регион
    area: string;
    // ИНН
    invoiceINN?: string;
    // Счет оплаты
    invoiceAccount?: string;
    // Клиент
    clientType: ClientType;
    // Одобрена
    approved: boolean;
    // Время одобрения
    approvedAt: Date;
    // Одобривший заявку
    approver: User;
};
