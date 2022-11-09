import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {PriceDirective} from "./price.directive";


@NgModule({
    declarations: [PriceDirective],
    imports: [CommonModule],
    exports: [PriceDirective],
})
export class PriceModule {
}
