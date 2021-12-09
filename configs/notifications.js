import whatElseToDoConfig from "../configs/whatElseToDo.js"
import store from "./configStore.js";
import inquirer from "inquirer";

import welcomeNotificationConfig from "../configs/welcome.js"
import orderConfirmationNotificationConfig from "../configs/orderConfirmation.js"
import timeOffRequestNotificationConfig from "../configs/timeOffRequest.js"

import sendNotification from "./sendNotification.js";

const notificationsPrompt = () => {
    const questions = [
        {
            name: "notification",
            type: "list",
            message: "Which notification would you like to send?",
            choices: [
                {
                    key: "welcome",
                    name: "Welcome Notification",
                    value: "C1XM9AWT2QM4A8PJVSY8KPGW4WWA"
                },
                {
                    key: "timeoff",
                    name: "Time Off Notification",
                    value: "2PY62BMJTQ41QDHTN5Q219ZH33B9"
                },
                {
                    key: "orderconfirmation",
                    name: "Order Confirmation Notification",
                    value: "J2VPFZ7F9VMTMTKAZWREHE54659A"
                },
            ],
            filter: (value) => {
                store.set("postData.event", value);
                return value
            },
        },
        ...welcomeNotificationConfig,
        ...timeOffRequestNotificationConfig,
        ...orderConfirmationNotificationConfig,
    ];

    inquirer.prompt(questions).then((answers) => {
        sendNotification()
        console.log('Email notification sent! Please check your inbox :) \n')
        whatElseToDoConfig()
    });
}

export default notificationsPrompt;