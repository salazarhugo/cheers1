import {AfterViewInit, Component, OnInit } from '@angular/core';
import {UserService} from "../../shared/data/services/user.service";
import {fadeAnimation} from "../../animtaions";

@Component({
    selector: 'app-core',
    templateUrl: './core.component.html',
    styleUrls: ['./core.component.sass'],
    animations: [fadeAnimation],
})
export class CoreComponent implements OnInit, AfterViewInit {

    constructor(
        public userService: UserService,
    ) {
    }

    ngOnInit(): void {
    }

    ngAfterViewInit(): void {
    }
}
