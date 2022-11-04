import {Component, Input, OnInit} from '@angular/core';
import {Activity} from "../../data/activity.model";

@Component({
    selector: 'app-activity-item',
    templateUrl: './activity-item.component.html',
    styleUrls: ['./activity-item.component.sass']
})
export class ActivityItemComponent implements OnInit {

    @Input() activity: Activity

    constructor() {
    }

    ngOnInit(): void {
    }

    onImgError(event: any) {
        event.target.src = 'assets/default_profile_picture.png';
    }
}
