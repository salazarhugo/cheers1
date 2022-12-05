import {Component, OnInit} from '@angular/core';
import {UntypedFormControl, UntypedFormGroup} from "@angular/forms";
import {UserService} from "../../data/services/user.service";
import {AuthService} from "../../data/services/auth.service";
import {AngularFireAuth} from "@angular/fire/compat/auth";
import {Router} from "@angular/router";
import {firstValueFrom} from "rxjs";

@Component({
    selector: 'app-finish-sign-up',
    templateUrl: './finish-sign-up.component.html',
    styleUrls: ['./finish-sign-up.component.css', '../sign-in/sign-in.component.sass'],
})
export class FinishSignUpComponent implements OnInit {

    userForm = new UntypedFormGroup({
        name: new UntypedFormControl(''),
        username: new UntypedFormControl(''),
    });
    errorMessage: string | null = null

    constructor(
        private userService: UserService,
        private authService: AuthService,
        private afAuth: AngularFireAuth,
        private router: Router,
    ) {
        this.getRedirectResult()
    }

    async getRedirectResult() {
        try {
            const result = await this.afAuth.getRedirectResult()
            const user = result.user
            if (user) {
                console.log('You have been successfully logged in!');
            }
        } catch (e) {
            console.log(e);
            await this.router.navigate(['sign-in'])
        }
    }

    ngOnInit(): void {
    }

    async onSubmit() {
        const userForm = this.userForm.value
        const afUser = await this.afAuth.currentUser

        if (!afUser || this.userForm.invalid)
            return

        const userReq = {
            name: userForm.name,
            email: afUser.email!,
            username: userForm.username,
            email_verified: afUser.emailVerified,
            picture: afUser.photoURL,
        }
        try {
            await firstValueFrom(this.userService.register(userReq))
            await this.router.navigate(['profile'])
        } catch (e: any) {
            this.errorMessage = e
            console.log("Failed to create User", e)
        }
    }
}
