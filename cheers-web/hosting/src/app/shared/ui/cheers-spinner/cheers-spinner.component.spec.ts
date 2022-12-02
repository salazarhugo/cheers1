import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CheersSpinnerComponent } from './cheers-spinner.component';

describe('CheersSpinnerComponent', () => {
  let component: CheersSpinnerComponent;
  let fixture: ComponentFixture<CheersSpinnerComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CheersSpinnerComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(CheersSpinnerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
