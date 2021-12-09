import inquirer from "inquirer";
import request from 'request';
import store from "./configStore.js";
import whatElseToDoConfig from "../configs/whatElseToDo.js"


const sendNotification = () => {
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
            }
        }
    );
};

export default sendNotification;