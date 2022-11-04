import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';

import {SettingBillingRoutingModule} from './setting-billing-routing.module';
import {SettingBillingComponent} from './setting-billing.component';


@NgModule({
    declarations: [SettingBillingComponent],
    imports: [
        CommonModule,
        SettingBillingRoutingModule
    ],
    exports: [SettingBillingComponent]
})
export class SettingBillingModule {
}
