import {Component, OnInit, ViewEncapsulation} from '@angular/core';
import {UntypedFormControl, UntypedFormGroup} from "@angular/forms";
import {AuthService} from "../../data/services/auth.service";
import {UserService} from "../../data/services/user.service";

@Component({
    selector: 'app-register',
    templateUrl: './register.component.html',
    styleUrls: ['./register.component.sass', '../sign-in/sign-in.component.sass'],
})
export class RegisterComponent implements OnInit {

    userForm = new UntypedFormGroup({
        email: new UntypedFormControl(''),
        name: new UntypedFormControl(''),
        username: new UntypedFormControl(''),
        password: new UntypedFormControl(''),
    });

    public signInLinkSent: boolean = false
    isLoading = false

    constructor(
        private userService: UserService,
        private authService: AuthService,
    ) {
    }

    ngOnInit(): void {
    }

    onSubmit() {
        this.isLoading = true
        const userForm = this.userForm.value
        this.authService.sendSignUpLinkToEmail(userForm["email"]).then(() => {
            this.signInLinkSent = true
            this.isLoading = false
        })
    }
}
