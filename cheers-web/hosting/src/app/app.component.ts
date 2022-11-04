import {Component} from '@angular/core';
import {AngularFireAuth} from "@angular/fire/compat/auth";
import {Router} from "@angular/router";

@Component({
    selector: 'app-root',
    templateUrl: './app.component.html',
    styleUrls: ['./app.component.css']
})
export class AppComponent {
    title = 'cheers';
    isLoading: boolean = false

    constructor(
        public auth: AngularFireAuth,
        private router: Router,
    ) {
        this.auth.isSignInWithEmailLink(window.location.href).then((success: boolean) => {
            if (!success) {
                return;
            }
            this.isLoading = true

            let email = window.localStorage.getItem('emailForSignIn');
            if (!email) {
                email = window.prompt('Please provide your email for confirmation');
            }
            this.auth.signInWithEmailLink(email!, window.location.href)
                .then(async (result) => {
                    window.localStorage.removeItem('emailForSignIn');
                    await this.router.navigate(['home'])
                    this.isLoading = false
                })
                .catch((error) => {
                    console.log(error)
                    this.isLoading = false
                });
        })
    }
}
