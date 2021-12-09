import store from "./configStore.js";

function sendNotification() {
    var headers = {
        'Authorization': 'Bearer dk_prod_FABZRCY7PJ494JGPK6G81XDFFKAY'
    };

    var options = {
        url: 'https://api.courier.com/send',
        method: 'POST',
        headers: headers,
        body: JSON.stringify(store.get('postData'))
    };

    request(options,
        (error, response, body) => {
            if (!error && response.statusCode == 200) {
                store.set('sentMessageId', JSON.parse(body).messageId)
                console.log(body);
                console.log(store.all);
            }
        });
};

const welcomeNotificationConfig = [
    {
        name: "welcome.companyName",
        type: "input",
        message: "Enter your company name:",
        when: (answers) => answers.notification === 'C1XM9AWT2QM4A8PJVSY8KPGW4WWA',
        filter: (value) => {
            store.set('postData.data', {
                ...store.get('postData.data'),
                "company_name": value,
            });
            return value
        },
    },
    {
        name: "welcome.companyTagline",
        type: "input",
        message: "Enter your company tagline:",
        when: (answers) => answers.notification === 'C1XM9AWT2QM4A8PJVSY8KPGW4WWA',
        filter: (value) => {
            store.set('postData.data', {
                ...store.get('postData.data'),
                "company_tagline": value,
            });
            return value
        },
    },
    {
        name: "welcome.companyDesc",
        type: "input",
        message: "Enter your company description:",
        when: (answers) => answers.notification === 'C1XM9AWT2QM4A8PJVSY8KPGW4WWA',
        filter: (value) => {
            store.set('postData.data', {
                ...store.get('postData.data'),
                "company_description": value,
            });
            return value
        },
    },
    {
        name: "email",
        type: "input",
        message: "What email would you like to send this to?",
        when: (answers) => answers.notification === 'C1XM9AWT2QM4A8PJVSY8KPGW4WWA',
        filter: (value) => {
            store.set('postData.profile', {
                ...store.get('postData.profile'),
                "email": value,
            });
            return value
        },
    },
    {
        name: "confirmation",
        type: 'confirm',
        message: "Ready to send a preview to yourself?",
        default: true,
        when: (answers) => answers.notification === 'C1XM9AWT2QM4A8PJVSY8KPGW4WWA',
        transformer: (value) => {
            // sendNotification();
            return "Welcome notification sent. Check your inbox."
        }
    }
];

export default welcomeNotificationConfig;