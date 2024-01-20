import {Injectable} from '@angular/core';
import {AngularFireAuth} from "@angular/fire/compat/auth";
import {GoogleAuthProvider} from 'firebase/auth';
import {UserService} from "./user.service";
import {ApiService} from "./api.service";
import {Router} from "@angular/router";
import {Credential, PublicKeyCredentialOptions} from "../models/webauthn.models";
import {map, Observable, of, switchMap} from "rxjs";

export interface StatusResponse {
    status: string;
    message?: string;
}

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

    toAllowCredentials(credentials: Credential[]): PublicKeyCredentialDescriptor[] {
        return credentials.map(credential => {
            // credential.id = atob(credential.id)
            // credential.publicKey = atob(credential.publicKey)
            return {
                id: Uint8Array.from(credential.publicKey, c => c.charCodeAt(0)).buffer,
                type: "public-key",//credential.attestationType as PublicKeyCredentialType,
                transports: ["internal"],//credential.transport as AuthenticatorTransport[],
            }
        })
    }

    loginUser(username: string): Observable<StatusResponse> {
        return this.api.beginLogin(username).pipe(
            switchMap(async response => {
                if (response == null)
                    return
                console.log(response)
                const getCredential = {
                    publicKey: {
                        user: {
                            id: response.userId,
                        },
                        timeout: response.timeout,
                        allowCredentials: this.toAllowCredentials(response.allowCredentials),
                        challenge: new TextEncoder().encode(response.challenge).buffer,
                        rpId: response.replyingPartyId,
                        userVerification: response.userVerification as UserVerificationRequirement,
                    },
                };
                return navigator.credentials.get(getCredential);
            }),
            map(() => {
                // this.api.finishLogin(result)
                const status: StatusResponse = {
                    status: "success",
                }
                return status
            })
        );
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
                await this.router.navigate(['parties'])
            }
        } catch (e) {
            console.log(e);
        }
    }

    async signInWithEmailAndPassword(email: string, password: string) {
        try {
            const res = await this.afAuth.signInWithEmailAndPassword(email, password)
            await this.router.navigate(['parties'])
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
