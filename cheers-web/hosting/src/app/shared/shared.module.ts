import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {MatRippleModule} from "@angular/material/core";
import {RouterModule} from "@angular/router";
import {FlexModule} from "@angular/flex-layout";
import {MaterialModule} from "../material/material.module";
import {CompleteComponent} from "./ui/complete/complete.component";
import {FinishSignUpComponent} from "./ui/finish-sign-up/finish-sign-up.component";
import {RegisterComponent} from "./ui/register/register.component";
import {SigninComponent} from "./ui/signin/signin.component";
import {SignInComponent} from "./ui/sign-in/sign-in.component";
import {TicketsComponent} from "./ui/tickets/tickets.component";
import {TicketingComponent} from "./ui/ticketing/ticketing.component";
import {LottieModule} from "ngx-lottie";
import {CoreRoutingModule} from "../core/core-routing.module";
import {DropzoneDirective} from './data/dropzone.directive';
import {CheersSpinnerModule} from "./ui/cheers-spinner/cheers-spinner.module";
import {CheersButtonComponent } from './ui/cheers-button/cheers-button.component';


const components = [
    CompleteComponent,
    FinishSignUpComponent,
    RegisterComponent,
    SigninComponent,
    SignInComponent,
    TicketsComponent,
    TicketingComponent,
    DropzoneDirective,
]

@NgModule({
    declarations: [
        ...components,
    ],
    imports: [
        CommonModule,
        MatRippleModule,
        FlexModule,
        MaterialModule,
        LottieModule,
        CoreRoutingModule,
        RouterModule,
        CheersSpinnerModule,
        CheersButtonComponent,
    ],
    exports: [
        ...components,
    ]
})
export class SharedModule {
}
