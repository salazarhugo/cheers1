import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';

import {CheersSpinnerRoutingModule} from './cheers-spinner-routing.module';
import {CheersSpinnerComponent} from './cheers-spinner.component';
import {MatLegacyProgressSpinnerModule as MatProgressSpinnerModule} from "@angular/material/legacy-progress-spinner";
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
