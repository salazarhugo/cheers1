import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import {BusinessPayoutsComponent} from "./business-payouts.component";

const routes: Routes = [{path: '', component: BusinessPayoutsComponent}];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class BusinessPayoutsRoutingModule { }
