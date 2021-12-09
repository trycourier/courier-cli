import Configstore from 'configstore';

// Create a Configstore instance.
const config = new Configstore('courier-cli');

const notificationsPostData = {
    event: "",
    recipient: "ffe1d347-74d1-4a44-8cb5-3823dd99974c",
    brand: "B7MPCPHABYME0VM1A56SJZWXXDT3",
    override: {},
    data: {},
    profile: {}
};

config.clear()
config.set('postData', notificationsPostData);

export default config;