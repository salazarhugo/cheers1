import {Injectable} from '@angular/core';
import {AngularFireAuth} from "@angular/fire/compat/auth";
import {GoogleAuthProvider} from 'firebase/auth';
import {UserService} from "./user.service";
import {ApiService} from "./api.service";
import {Router} from "@angular/router";

@Injectable({
    providedIn: 'root'
})
export class AuthService {

    constructor(
        private api: ApiService,
        private afAuth: AngularFireAuth,
        private userService: UserService,
        private router: Router,
    ) {
        this.afAuth.currentUser.then(user => {
            if (user)
                localStorage.setItem('user', JSON.stringify(user));
        })

        this.afAuth.authState.subscribe((user) => {
            if (user) {
                localStorage.setItem('user', JSON.stringify(user));
            } else {
                localStorage.setItem('user', 'null');
            }
        });
    }

    createUserWithEmailAndPassword(email: string, password: string) {
        return new Promise<any>((resolve, reject) => {
            this.afAuth.createUserWithEmailAndPassword(email, password)
                .then(res => {
                    resolve(res);
                }, err => reject(err))
        })
    }

    async signInWithGoogle() {
        try {
            const result = await this.afAuth.signInWithPopup(new GoogleAuthProvider())
            if (result.user) {
                console.log('You have been successfully logged in!');
                await this.router.navigate(['profile'])
            }
        } catch (e) {
            console.log(e);
        }
    }

    async signInWithEmailAndPassword(email: string, password: string) {
        try {
            const res = await this.afAuth.signInWithEmailAndPassword(email, password)
            await this.router.navigate(['profile'])
        } catch (e) {
            console.log(e);
        }
    }

    sendSignInLinkToEmail(email: string) {
        const actionCodeSettings = {
            url: window.location.origin + "/sign-in",
            handleCodeInApp: true,
            iOS: {
                bundleId: "com.salazar.cheers"
            },
            android: {
                packageName: "com.salazar.cheers",
                installApp: true,
                minimumVersion: '12'
            },
            dynamicLinkDomain: 'cheers2cheers.page.link'
        };

        return this.afAuth.sendSignInLinkToEmail(email, actionCodeSettings)
            .then(() => {
                window.localStorage.setItem('emailForSignIn', email);
            })
            .catch((error) => {
                const errorCode = error.code;
                const errorMessage = error.message;
            });
    }

    sendSignUpLinkToEmail(email: string) {
        const actionCodeSettings = {
            url: window.location.origin + "/finish-sign-up",
            // url: "https://web-cheers.web.app/finish-sign-up",
            handleCodeInApp: true,
            iOS: {
                bundleId: "com.salazar.cheers"
            },
            android: {
                packageName: "com.salazar.cheers",
                installApp: true,
                minimumVersion: '12'
            },
            dynamicLinkDomain: 'cheers2cheers.page.link'
        };

        return this.afAuth.sendSignInLinkToEmail(email, actionCodeSettings)
            .then(() => {
                window.localStorage.setItem('emailForSignIn', email);
            })
            .catch((error) => {
                const errorCode = error.code;
                const errorMessage = error.message;
            });
    }

    verifyUser(uid: string) {
        return this.api.verifyUser(uid)
    }

    promoteToBusiness(uid: string) {
        return this.api.promoteToBusiness(uid)
    }

    signOut() {
        return this.afAuth.signOut();
    }

}
