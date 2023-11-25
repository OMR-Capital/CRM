import { FakePurchasesApi, PurchasesApi } from './purchases';

export class Client {
    static readonly BASE_URL = '/';

    makeUrl(path: string, params?: Record<string, string>): string {
        return `${Client.BASE_URL}/${path}?${new URLSearchParams(params)}`;
    }

    purchasesApi(): PurchasesApi {
        return new PurchasesApi(this);
    }
}

export class FakeClient extends Client {
    purchasesApi(): PurchasesApi {
        return new FakePurchasesApi(this);
    }
}
