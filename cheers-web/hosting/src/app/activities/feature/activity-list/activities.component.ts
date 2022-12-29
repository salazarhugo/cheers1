import {Component, OnInit} from '@angular/core';
import {Activity} from "../../data/activity.model";
import {ActivityService} from "../../data/activity.service";
import {firstValueFrom} from "rxjs";
import {UserService} from "../../../shared/data/services/user.service";

@Component({
    selector: 'app-activity-list',
    templateUrl: './activities.component.html',
    styleUrls: ['./activities.component.sass']
})
export class ActivitiesComponent implements OnInit {

    activities: Activity[] = []

    constructor(
        private userService: UserService,
        private activityService: ActivityService,
    ) {
    }

    async ngOnInit() {
        const activities = await firstValueFrom(this.activityService.listActivity())
        this.activities = activities
    }

}
