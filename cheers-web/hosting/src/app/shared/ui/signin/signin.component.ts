import {Component, OnInit} from '@angular/core';
import {UntypedFormControl, UntypedFormGroup} from "@angular/forms";
import {UserService} from "../../data/services/user.service";
import {AuthService} from "../../data/services/auth.service";
import {ActivatedRoute, Router} from "@angular/router";
import {AngularFireAuth} from "@angular/fire/compat/auth";
import { FirebaseError } from '@angular/fire/app/firebase';

@Component({
    selector: 'app-signin',
    templateUrl: './signin.component.html',
    styleUrls: ['./signin.component.sass', '../sign-in/sign-in.component.sass']
})
export class SigninComponent implements OnInit {

    userForm = new UntypedFormGroup({
        email: new UntypedFormControl(''),
        password: new UntypedFormControl(''),
    });

    withPassword: boolean = false
    signInLinkSent: boolean = false
    isLoading: boolean = false
    errorMessage: string = ""

    constructor(
        private userService: UserService,
        private authService: AuthService,
        private router: Router,
        private afAuth: AngularFireAuth,
    ) {
    }

    ngOnInit(): void {
    }

    signInWithGoogle() {
        this.authService.signInWithGoogle()
    }

    onSignIn() {
        this.isLoading = true
        this.errorMessage = ""
        const userForm = this.userForm.value

        if (this.withPassword) {
            this.authService.signInWithEmailAndPassword(
                userForm["email"],
                userForm["password"],
            ).then(res => {
                this.isLoading = false
            }).catch((err: FirebaseError) => {
                this.errorMessage = err.message
                this.isLoading = false
            })
        } else {
            this.authService.sendSignInLinkToEmail(userForm["email"]).then(() => {
                this.signInLinkSent = true
                this.isLoading = false
            })
        }
    }
}
