require_relative 'boot'

require 'rails/all'
require_relative('../app/models/article.rb')

# Require the gems listed in Gemfile, including any gems
# you've limited to :test, :development, or :production.
Bundler.require(*Rails.groups)

module Planetocd
  class Application < Rails::Application

    attr_accessor :articles
    @articles

    def load_articles(path)
      @articles = Hash.new
      to_convert = Dir.glob("#{path}/*.mdocd")
      to_convert.each do |article_path|
        article = Article.new(article_path)
        if @articles[article.language] == nil
          @articles[article.language] = Hash.new
        end
        @articles[article.language][article.id] = article
      end
    end

    # Initialize configuration defaults for originally generated Rails version.
    config.load_defaults 5.2

    config.exceptions_app = self.routes
    config.middleware.insert_before(Rack::Sendfile, Rack::Deflater)
    config.assets.compress = true

    # Settings in config/environments/* take precedence over those specified here.
    # Application configuration can go into files in config/initializers
    # -- all .rb files in that directory are automatically loaded after loading
    # the framework and any gems in your application.
    
    load_articles(Rails.root.join('app', 'assets', 'articles'))
  end
end
