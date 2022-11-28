import {Component, OnInit, ViewChild} from '@angular/core';
import {OrderService} from "../../../shared/data/services/order.service";
import {Order} from "../../../../gen/ts/cheers/order/v1/order_service";
import {MatSort} from "@angular/material/sort";
import {MatTableDataSource} from "@angular/material/table";
import {MatSnackBar} from "@angular/material/snack-bar";
import {Clipboard} from '@angular/cdk/clipboard';
import {PaymentService} from "../../../payments/data/payment.service";

@Component({
    selector: 'app-business-orders',
    templateUrl: './business-orders.component.html',
    styleUrls: ['./business-orders.component.sass']
})
export class BusinessOrdersComponent implements OnInit {
    displayedColumns: string[] = ['status', 'date', 'amount', 'party', 'firstName', 'lastName', 'email', 'more'];
    dataSource: MatTableDataSource<Order> = new MatTableDataSource();

    constructor(
        private clipboard: Clipboard,
        private snackBar: MatSnackBar,
        private orderService: OrderService,
        private paymentService: PaymentService,
    ) {
    }

    @ViewChild(MatSort) sort: MatSort;

    ngAfterViewInit() {
        this.dataSource.sort = this.sort;
    }

    async ngOnInit() {
        const res = await this.orderService.getCurrentUserOrders()
        res.subscribe(data => {
            this.dataSource.data = data
        })
    }

    /** Copy order ID to clipboard */
    copyOrderID(orderID: string) {
        this.clipboard.copy(orderID);
        this.snackBar.open("Copied order id", "Hide", {
            duration: 3000
        })
    }

    /** Refund payment */
    refundPayment(paymentID: string) {
        this.paymentService.refundPayment(paymentID).subscribe()
    }

    /** Gets the total amount of all orders. */
    getTotalCost() {
        return this.dataSource.data.map(t => t.amount).reduce((acc, value) => acc + value, 0);
    }
}