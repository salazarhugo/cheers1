import {Injectable} from '@angular/core';

@Injectable({
    providedIn: 'root'
})
export class ActivityService {

    // private client: ActivityServiceClient;

    constructor() {
        // this.client = new ActivityServiceClient(
        //     'https://activity-service-r3a2dr4u4a-nw.a.run.app:443', {
        //         debug: true
        //     });
    }

    listActivity(userId: string): void {
        // const request = new ListActivityRequest();
        // request.setUserId(userId);
        // this.client.listActivity(
        //     request, (error: any, response: ListActivityResponse | null) => {
        //         // Your code to handle error & response.
        //         console.log('Error: ' + error);
        //         console.log('ListActivityResponse: ' + response);
        //     });
    }
}
