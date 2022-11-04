import {Component, OnInit} from '@angular/core';
import {Activity} from "../../data/activity.model";

@Component({
    selector: 'app-activity-list',
    templateUrl: './activities.component.html',
    styleUrls: ['./activities.component.sass']
})
export class ActivitiesComponent implements OnInit {

    activities = [
        new Activity(),
        new Activity(),
        new Activity(),
        new Activity(),
    ]
    constructor() {
    }

    ngOnInit(): void {
    }

}
