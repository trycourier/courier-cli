import store from "./configStore.js";


const orderConfirmationNotificationConfig = [
    {
        name: "orderconfirmation.companyName",
        type: "input",
        message: "Enter your company name:",
        when: (answers) => answers.notification === 'J2VPFZ7F9VMTMTKAZWREHE54659A',
        filter: (value) => {
            store.set('postData.data', {
                ...store.get('postData.data'),
                "company_name": value,
            });
            return value
        },
    },
    {
        name: "email",
        type: "input",
        message: "What's your email?",
        when: (answers) => answers.notification === 'J2VPFZ7F9VMTMTKAZWREHE54659A',
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
        when: (answers) => answers.notification === 'J2VPFZ7F9VMTMTKAZWREHE54659A',
    },
];

export default orderConfirmationNotificationConfig;