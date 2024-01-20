import {Component, OnInit} from '@angular/core';
import {UntypedFormControl, UntypedFormGroup} from "@angular/forms";
import {UserService} from "../../data/services/user.service";
import {AuthService} from "../../data/services/auth.service";
import {ActivatedRoute, Router} from "@angular/router";
import {AngularFireAuth} from "@angular/fire/compat/auth";
import { FirebaseError } from '@angular/fire/app/firebase';
import {firstValueFrom} from "rxjs";
import {MatSnackBar} from "@angular/material/snack-bar";

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
        private _snackBar: MatSnackBar,
    ) {
    }

    ngOnInit(): void {
        if (window.PublicKeyCredential) {
            PublicKeyCredential.isUserVerifyingPlatformAuthenticatorAvailable()
                .then(uvpaEnabled => {
                    if (uvpaEnabled) {
                        console.log("uvpa enabled")
                    } else {
                        console.log("uvpa not enabled")
                    }
                });
        }
    }

    openSnackBar(message: string) {
        this._snackBar.open(message);
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
        const username = userForm["username"]

        this.authService.loginUser(username).subscribe({
            next: () => {},
            error: (err) => {
                this.openSnackBar(err)
                this.isLoading = false
            }
        })
    }
}

