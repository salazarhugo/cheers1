import {Component, OnInit} from '@angular/core';
import {AngularFireAuth} from "@angular/fire/compat/auth";
import {Router} from "@angular/router";
import {ThemeService} from "./core/data/theme.service";
import {Observable} from "rxjs";
import { getMessaging, getToken, onMessage } from "firebase/messaging";
import {environment} from "../environments/environment";
import {MatSnackBar} from "@angular/material/snack-bar";
import {NotificationService} from "./core/data/notification.service";

@Component({
    selector: 'app-root',
    templateUrl: './app.component.html',
    styleUrls: ['./app.component.sass']
})
export class AppComponent implements OnInit {
    title = 'cheers';
    isLoading: boolean = false
    isDarkTheme: Observable<boolean>;

    constructor(
        public auth: AngularFireAuth,
        private router: Router,
        public themeService: ThemeService,
        private _snackBar: MatSnackBar,
        private notificationService: NotificationService,
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

    ngOnInit(): void {
        this.isDarkTheme = this.themeService.isDarkTheme;
        this.requestPermission();
        this.listen();
    }

    openSnackBar(message: string) {
        this._snackBar.open(message, undefined, {
           horizontalPosition: "start",
            verticalPosition: "bottom",
            politeness: "assertive",
        });
    }

    requestPermission() {
        const messaging = getMessaging();
        getToken(messaging,
            { vapidKey: environment.firebaseConfig.vapidKey}).then(
            (currentToken) => {
                if (currentToken) {
                    this.notificationService.createToken(currentToken).subscribe(res => console.log(res))
                } else {
                    console.log('No registration token available. Request permission to generate one.');
                }
            }).catch((err) => {
            console.log('An error occurred while retrieving token. ', err);
        });
    }

    listen() {
        const messaging = getMessaging();
        onMessage(messaging, (payload) => {
            console.log('Message received. ', payload);
            const message = payload?.data
            if (message == undefined)
                return
            this.openSnackBar(message["title"] + "\n" + message["body"])
        });
    }
}
