import {Component, Input, OnInit} from '@angular/core';

@Component({
    selector: 'app-user-chip',
    templateUrl: './user-chip.component.html',
    styleUrls: ['./user-chip.component.sass']
})
export class UserChipComponent implements OnInit {

    @Input() user: any

    constructor() {
    }

    ngOnInit(): void {
    }

    onImgError(event: any) {
        event.target.src = 'assets/default_profile_picture.png';
    }
}
