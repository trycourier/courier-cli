#!/usr/bin/env node

import chalk from "chalk";
import boxen from "boxen";
import clear from "clear";
import figlet from "figlet";
import inquirer from "inquirer";

import randomJoke from "../configs/randomJoke.js"
import notificationsConfig from "../configs/notifications.js"


const main = async () => {
    clear();

    const banner = chalk.yellow(
        figlet.textSync('Courier CLI', { horizontalLayout: 'full' })
    )
    const boxenOptions = {
        padding: 1,
        margin: 1,
        borderStyle: "round",
        borderColor: "#C705E9",
        backgroundColor: "#555555"
    };
    console.log(boxen(banner, boxenOptions));

    const questions = [
        {
            name: "whatToDo",
            type: "list",
            message: "What do you want to do?",
            choices: [
                {
                    key: "sendNotification",
                    name: "Send a Notification",
                    value: "sendNotification"
                },
                new inquirer.Separator(),
                {
                    key: "randomJoke",
                    name: "Tell me a joke",
                    value: "randomJoke"
                },
                {
                    name: 'Contact support',
                    disabled: 'Unavailable at this time',
                },
            ]
        },
    ];

    inquirer.prompt(questions).then((answers) => {
        switch (answers.whatToDo) {
            case "sendNotification":
                notificationsConfig();
                break;
            case "randomJoke":
                randomJoke().then((joke) => console.log(joke))
                setTimeout(() => {
                    clear()
                    main()    
                }, 10000)
                break;

            default:
                break;
        }
    });
};

await main();

export default main;