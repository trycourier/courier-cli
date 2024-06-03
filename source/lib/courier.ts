import { CourierClient } from "@trycourier/courier";

const getClient = (environment: 'test' | 'production' = 'production') => {

    return new CourierClient({
        authorizationToken: environment ? process.env['COURIER_AUTH_TOKEN'] : process.env['COURIER_AUTH_TOKEN_TEST'] ,
    });
}

export default getClient;