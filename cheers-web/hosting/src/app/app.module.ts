import {NgModule} from '@angular/core';
import {BrowserModule} from '@angular/platform-browser';
import {AppComponent} from './app.component';
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import {HttpClientModule} from "@angular/common/http";
import {httpInterceptorProviders} from "./shared/data/interceptors";
import {AngularFireModule, FIREBASE_OPTIONS} from "@angular/fire/compat";
import {environment} from "../environments/environment.prod";
import {AngularFireAuthModule} from "@angular/fire/compat/auth";
import {RouterModule} from "@angular/router";
import {AppRoutingModule} from "./app-routing.module";
import {GrpcCoreModule} from "@ngx-grpc/core";
import {GrpcWebClientModule} from "@ngx-grpc/grpc-web-client";
import {OverlayModule} from "@angular/cdk/overlay";
import {LottieModule} from "ngx-lottie";
import {getApp, initializeApp, provideFirebaseApp} from "@angular/fire/app";
import {initializeAppCheck, provideAppCheck, ReCaptchaV3Provider} from '@angular/fire/app-check';
import {MaterialModule} from "./material/material.module";
import {SharedModule} from "./shared/shared.module";
import {HomeModule} from "./home/home.module";
import {PaymentsModule} from "./payments/payments.module";
import {StoriesModule} from "./stories/stories.module";
import {CoreModule} from "./core/core.module";
import {PartyShellModule} from "./parties/features/party-shell/party-shell.module";
import {MAT_SNACK_BAR_DEFAULT_OPTIONS} from "@angular/material/snack-bar";
import {ManagePartyModule} from "./manage-parties/feature/manage-party/manage-party.module";
import { RefundPaymentDialogComponent } from './business/ui/refund-payment-dialog/refund-payment-dialog.component';
import {MapModule} from "./map/map.module";
import {PostCreateDialogModule} from "./posts/ui/post-create-dialog/post-create-dialog.module";
import {TopbarModule} from "./home/ui/topbar/topbar.module";


export function playerFactory(): any {
    return import("lottie-web");
}

if (!environment.production)
    (<any>window).FIREBASE_APPCHECK_DEBUG_TOKEN = true;

@NgModule({
    declarations: [
        AppComponent,
        RefundPaymentDialogComponent,
    ],
    imports: [
        BrowserModule,
        BrowserAnimationsModule,
        AppRoutingModule,
        HttpClientModule,
        AngularFireAuthModule,
        OverlayModule,
        MaterialModule,
        RouterModule,
        LottieModule.forRoot({player: playerFactory}),
        AngularFireModule.initializeApp(environment.firebaseConfig),
        provideFirebaseApp(() => initializeApp(environment.firebaseConfig)),
        provideAppCheck(() => initializeAppCheck(getApp(), {
            provider: new ReCaptchaV3Provider(environment.recaptchaSiteKey),
            isTokenAutoRefreshEnabled: true
        })),
        SharedModule,
        HomeModule,
        PaymentsModule,
        StoriesModule,
        CoreModule,
        ManagePartyModule,
        PartyShellModule,
        MapModule,
        PostCreateDialogModule,
        TopbarModule,
    ],
    providers: [
        httpInterceptorProviders,
        {provide: FIREBASE_OPTIONS, useValue: environment.firebaseConfig},
        {provide: MAT_SNACK_BAR_DEFAULT_OPTIONS, useValue: {duration: 2500}}
    ],
    exports: [],
    bootstrap: [AppComponent]
})
export class AppModule {

}
