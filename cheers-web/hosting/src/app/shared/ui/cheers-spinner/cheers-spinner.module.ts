import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';

import {CheersSpinnerRoutingModule} from './cheers-spinner-routing.module';
import {CheersSpinnerComponent} from './cheers-spinner.component';
import {MatProgressSpinnerModule} from "@angular/material/progress-spinner";
import {FlexModule} from "@angular/flex-layout";


@NgModule({
    declarations: [
        CheersSpinnerComponent
    ],
    imports: [
        CommonModule,
        CheersSpinnerRoutingModule,
        MatProgressSpinnerModule,
        FlexModule
    ],
    exports: [
        CheersSpinnerComponent
    ],
})
export class CheersSpinnerModule {
}
