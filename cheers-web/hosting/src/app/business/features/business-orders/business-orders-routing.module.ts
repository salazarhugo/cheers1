import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import {BusinessOrdersComponent} from "./business-orders.component";

const routes: Routes = [{path: '', component: BusinessOrdersComponent}];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class BusinessOrdersRoutingModule { }
