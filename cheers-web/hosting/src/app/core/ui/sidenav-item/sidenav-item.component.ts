import {Component, Input, OnInit} from '@angular/core';
import {Router} from "@angular/router";

@Component({
    selector: 'app-sidenav-item',
    templateUrl: './sidenav-item.component.html',
    styleUrls: ['./sidenav-item.component.sass']
})
export class SidenavItemComponent implements OnInit {

    @Input() item: any
    @Input() business: boolean

    constructor(
        public router: Router,
    ) {
    }

    ngOnInit(): void {
    }

}
