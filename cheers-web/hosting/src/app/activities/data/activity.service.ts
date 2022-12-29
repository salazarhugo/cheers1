import {Injectable} from '@angular/core';
import {map, Observable} from "rxjs";
import {Activity, toActivity} from "./activity.model";
import {HttpClient} from "@angular/common/http";
import {ListActivityResponse} from "../../../gen/ts/cheers/activity/v1/activity_service";
import {environment} from "../../../environments/environment";

@Injectable({
    providedIn: 'root'
})
export class ActivityService {

    constructor(
        private http: HttpClient,
    ) {
    }

    listActivity(): Observable<Activity[]> {
        return this.http.get<ListActivityResponse>(`${environment.GATEWAY_URL}/v1/activities`)
            .pipe(map(res => res.activities.map(activity => toActivity(activity))))
    }
}
