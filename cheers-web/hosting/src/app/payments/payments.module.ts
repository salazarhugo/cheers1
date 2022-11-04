import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {PaymentComponent} from "./features/payment/payment.component";
import {ElementsComponent} from "./ui/elements/elements.component";
import {MaterialModule} from "../material/material.module";
import {RouterModule} from "@angular/router";


const components = [
    PaymentComponent,
    ElementsComponent,
]

@NgModule({
    declarations: [
        ...components
    ],
    imports: [
        CommonModule,
        MaterialModule,
        RouterModule,
    ],
    exports: [
        ...components
    ],
})
export class PaymentsModule {
}
