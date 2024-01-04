import {Component, OnInit} from '@angular/core';
import {UntypedFormControl, UntypedFormGroup} from "@angular/forms";
import {UserService} from "../../data/services/user.service";
import {AuthService} from "../../data/services/auth.service";
import {ActivatedRoute, Router} from "@angular/router";
import {AngularFireAuth} from "@angular/fire/compat/auth";
import { FirebaseError } from '@angular/fire/app/firebase';
import {firstValueFrom} from "rxjs";

@Component({
    selector: 'app-signin',
    templateUrl: './signin.component.html',
    styleUrls: ['./signin.component.sass', '../sign-in/sign-in.component.sass']
})
export class SigninComponent implements OnInit {

    userForm = new UntypedFormGroup({
        username: new UntypedFormControl(''),
    });

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

    async signInWithGoogle() {
        this.isLoading = true
        await this.authService.signInWithGoogle()
        this.isLoading = false
    }

    async onSignIn() {
        this.isLoading = true
        this.errorMessage = ""
        const userForm = this.userForm.value

        const result = await firstValueFrom(this.authService.loginUser(userForm["username"]))
        this.isLoading = false
    }
}
