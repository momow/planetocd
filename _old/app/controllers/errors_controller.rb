class ErrorsController < BaseController
    def internal_error
        render 'internal_error', :status => 500
    end
end
