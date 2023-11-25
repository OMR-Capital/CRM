import { TEST_PURCHASES } from '$lib/testData';
import type { Client } from './client';
import type { Purchase } from './models';

export class PurchasesApi {
    constructor(private client: Client) {}

    async getPurchases(): Promise<Purchase[]> {
        const response = await fetch(this.client.makeUrl('/purchases'));
        return (await response.json()) as Purchase[];
    }
}

export class FakePurchasesApi extends PurchasesApi {
    async getPurchases(): Promise<Purchase[]> {
        return TEST_PURCHASES;
    }
}
