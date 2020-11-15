class CreateLikes < ActiveRecord::Migration[5.2]
  def change
    create_table :likes do |t|
      t.string :ip_address,   null: false
      t.string :user_agent,   null: false
      t.integer :article_id,  null: false

      t.timestamps
    end
  end
end
