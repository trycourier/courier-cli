import store from "./configStore.js";


const timeOffRequestNotificationConfig = [
    {
        name: "timeoff.name",
        type: "input",
        message: "What's your name?",
        when: (answers) => answers.notification === '2PY62BMJTQ41QDHTN5Q219ZH33B9',
        filter: (value) => {
            store.set('postData.data', {
                ...store.get('postData.data'),
                "name": value,
            });
            return value
        },
    },
    {
        name: "timeoff.startdate",
        type: "input",
        message: "When will your vacation start?",
        when: (answers) => answers.notification === '2PY62BMJTQ41QDHTN5Q219ZH33B9',
        filter: (value) => {
            store.set('postData.data', {
                ...store.get('postData.data'),
                "start_date": value,
            });
            return value
        },
    },
    {
        name: "timeoff.enddate",
        type: "input",
        message: "When will you return to work?",
        when: (answers) => answers.notification === '2PY62BMJTQ41QDHTN5Q219ZH33B9',
        filter: (value) => {
            store.set('postData.data', {
                ...store.get('postData.data'),
                "end_date": value,
            });
            return value
        },
    },
    {
        name: "timeoff.reason",
        type: "input",
        message: "Why are you taking time off?",
        when: (answers) => answers.notification === '2PY62BMJTQ41QDHTN5Q219ZH33B9',
        filter: (value) => {
            store.set('postData.data', {
                ...store.get('postData.data'),
                "reason": value,
            });
            return value
        },
    },
    {
        name: "email",
        type: "input",
        message: "What's your email?",
        when: (answers) => answers.notification === '2PY62BMJTQ41QDHTN5Q219ZH33B9',
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
        when: (answers) => answers.notification === '2PY62BMJTQ41QDHTN5Q219ZH33B9',
    },
];

export default timeOffRequestNotificationConfig;