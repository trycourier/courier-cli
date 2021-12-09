import store from "./configStore.js";
import inquirer from "inquirer";
import main from "../bin/index.js"


const whatElseToDoPrompt = () => {
    const whatElseToDoQuestions = [
        {
            name: "whatElseToDo",
            type: "list",
            message: "Do you want to do anything else?",
            choices: [
                // new inquirer.Separator(),
                {
                    key: "goAgain",
                    name: "Give it another go?",
                    value: "goAgain"
                },
                {
                    key: "viewLogs",
                    name: "View logs",
                    value: "viewLogs"
                },
                {
                    key: "inviteDesigner",
                    name: "Invite a designer to customize your template",
                    value: "inviteDesigner"
                },
                {
                    key: "integrateEmailProvider",
                    name: "Integrate your own email provider",
                    value: "integrateEmailProvider"
                },
                {
                    key: "courierDocs",
                    name: "Explore Courier docs",
                    value: "courierDocs"
                },
                {
                    key: "contactSupport",
                    name: 'Contact support',
                    disabled: 'Unavailable at this time',
                    value: "contactSupport",
                },
                {
                    key: "chatOnDiscord",
                    name: "Chat now with a Courier expert in Discord",
                    value: "chatOnDiscord"
                },
                {
                    key: "pricingPlans",
                    name: "See pricing plans",
                    value: "pricingPlans"
                },
                {
                    key: "howTosOnYoutube",
                    name: "Watch how-to videos on our Youtube channel",
                    value: "howTosOnYoutube"
                },
            ],
        },
    ];

    inquirer.prompt(whatElseToDoQuestions).then((answers) => {
        switch (answers.whatElseToDo) {
            case "goAgain":
                 main();
                
                break;

            case "viewLogs":
                const url = `https://app.courier.com/data/messages?message=${store.get('sentMessageId')}`;
                console.log(url);
                break;

            default:
                break;
        }
    });
}


export default whatElseToDoPrompt;