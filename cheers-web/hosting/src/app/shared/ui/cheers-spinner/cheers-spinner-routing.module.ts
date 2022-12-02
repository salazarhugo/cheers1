import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import {CheersSpinnerComponent} from "./cheers-spinner.component";

const routes: Routes = [{path: '', component: CheersSpinnerComponent}];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class CheersSpinnerRoutingModule { }
