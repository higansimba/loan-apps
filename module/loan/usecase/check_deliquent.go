package usecase

import "context"

// CheckDeliquentInstallment implements LoanUsecase.
func (l *LoanUsecaseImpl) CheckDeliquentInstallment(ctx context.Context, loanid string) (int, error) {
	overduePayments, err := l.repo.GetOverduePayments(ctx, loanid)
	if err != nil {
		return 0, err
	}

	return len(overduePayments), nil

}
