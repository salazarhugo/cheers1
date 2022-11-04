import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import {SettingBillingComponent} from "./setting-billing.component";

const routes: Routes = [ { path: '', component: SettingBillingComponent }];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class SettingBillingRoutingModule { }
