import {Component, OnInit} from '@angular/core';


@Component({
    selector: 'app-sign-in',
    templateUrl: './sign-in.component.html',
    styleUrls: ['./sign-in.component.sass'],
})
export class SignInComponent implements OnInit {

    registerPage: boolean = false

    constructor() {
    }

    ngOnInit(): void {
    }

    onRegisterToggle() {
        this.registerPage = !this.registerPage
    }

}
